import Vue from 'vue'
import Router from 'vue-router'
import menus from '@/config/menu-config'

Vue.use(Router)

var routes = []
var children = []

menus.forEach((item) => {
  item.sub.forEach((sub) => {
    children.push({
      path: `/${sub.showPath}`,
      component: () => import(`@/views/${sub.componentName}`)
    })
  })
})

routes.push({
  path: `/`,
  component: () => import(`@/views/layout/Layout`),
  children: children
})

export default new Router({ routes })

// export const constantRouterMap = [
//   {
//     path: '/',
//     component: Layout,
//     children: [
//       {
//         // 当 /user/:id/profile 匹配成功，
//         // UserProfile 会被渲染在 User 的 <router-view> 中
//         path: '/',
//         component: MainPage
//       },
//       {
//         // 当 /user/:id/profile 匹配成功，
//         // UserProfile 会被渲染在 User 的 <router-view> 中
//         path: 'search',
//         component: search
//       }]
//   },
//   {
//     path: '*',
//     redirect: '/'
//   }
// ]

// export default new Router({
//   routes: constantRouterMap
// })
