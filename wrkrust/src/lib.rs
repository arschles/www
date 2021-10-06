use worker::*;

mod utils;

fn log_request(req: &Request) {
    console_log!(
        "{} - [{}], located at: {:?}, within: {}",
        Date::now().to_string(),
        req.path(),
        req.cf().coordinates().unwrap_or_default(),
        req.cf().region().unwrap_or("unknown region".into())
    );
}

async fn handle_redir<D>(req: Request, ctx: RouteContext<D>) -> Result<Response> {
    log_request(&req);
    let kv = ctx.kv("redirects")?;
    let resp: Result<Response> = match ctx.param("redir_path") {
        Some(redir_path) => {
            match kv.get(redir_path).await? {
                Some(redir_target) => {
                    let mut hdrs = Headers::new();
                    hdrs.set("Location", &redir_target.as_string())?;
                    Response::empty().map(|resp| {
                        resp.with_status(301).with_headers(hdrs)
                    })
                }
                None => {
                    Response::error("not found", 404)
                }
            }
        }
        None => {
            Response::error("not found", 404)
        }
    };
    resp
}

#[event(fetch)]
pub async fn main(req: Request, env: Env) -> Result<Response> {
    log_request(&req);

    // Optionally, get more helpful error messages written to the console in the case of a panic.
    utils::set_panic_hook();

    // Optionally, use the Router to handle matching endpoints, use ":name" placeholders, or "*name"
    // catch-alls to match on specific patterns. Alternatively, use `Router::with_data(D)` to
    // provide arbitrary data that will be accessible in each route via the `ctx.data()` method.
    let router: Router<()> = Router::new();

    // Add as many routes as your Worker needs! Each route will get a `Request` for handling HTTP
    // functionality and a `RouteContext` which you can use to  and get route parameters and
    // Environment bindings like KV Stores, Durable Objects, Secrets, and Variables.
    router
        .get_async("/:redir_path", handle_redir)
        .get("/worker-version", |_, ctx| {
            let version = ctx.var("WORKERS_RS_VERSION")?.to_string();
            Response::ok(version)
        })
        .run(req, env)
        .await
}
