import { getRouter } from './handler'

const router = getRouter()

addEventListener('fetch', (event) => {
  event.respondWith(router.handle(event.request, event))
})
