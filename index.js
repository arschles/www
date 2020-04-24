// resp = Response(request.args.get('challenge'))
// resp.headers['Content-Type'] = 'text/plain'
// resp.headers['X-Content-Type-Options'] = 'nosniff'


const Router = require('./router')

/**
 * Example of how router can be used in an application
 *  */
addEventListener('fetch', event => {
    event.respondWith(handleRequest(event.request))
})

function handler(request) {
    const init = {
        headers: { 'content-type': 'application/json' },
    }
    const body = JSON.stringify({ some: 'json' })
    return new Response(body, init)
}

async function handleRequest(request) {
    const r = new Router()
    r.get('/webhook', request => {
        challenge = request.headers.get("challenge")
        resp = new Response()
        resp.headers['Content-Type'] = 'text/plain'
        resp.headers['X-Content-Type-Options'] = 'nosniff'
        return resp
    })
    // Replace with the approriate paths and handlers
    // r.get('.*/bar', () => new Response('responding for /bar'))
    // r.get('.*/foo', request => handler(request))
    // r.post('.*/foo.*', request => handler(request))
    // r.get('/demos/router/foo', request => fetch(request)) // return the response from the origin

    r.get("/list_docs", req => {
        
    })
    r.get('/', () => new Response('Hello worker!')) // return a default message for the root route

    const resp = await r.route(request)
    return resp
}