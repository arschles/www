import {Router, Request, Obj} from 'itty-router'
// import * as redirects from '.'

declare const redirects: KVNamespace;

export function getRouter() : Router<void> {
  const router = Router()
  router.get('/*', async(req: Request): Promise<Response> => {
    try {
      const url = new URL(req.url)
      const path = url.pathname.slice(1)
      const redir_val = await redirects.get(path)
      if(redir_val == null) {
        return error(500, `no redirect for /${path}`)
      }
      return Response.redirect(redir_val)
    } catch(e: any) {
      return error(500, `exception thrown: ${e}`)
    }
  })

  return router
}

function error(code: number, body: string): Response {
  return new Response(body, {status: code});
}
