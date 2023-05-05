import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [{
  path: '/login',
  component: () => import('@/views/login/index'),
  hidden: true
},

{
  path: '/404',
  component: () => import('@/views/404'),
  hidden: true
},

{
  path: '/',
  component: Layout,
  redirect: '/realestate',
  children: [{
    path: 'realestate',
    name: 'Realestate',
    component: () => import('@/views/realestate/list/index'),
    meta: {
      title: '主页',
      icon: 'realestate'
    }
  }]
}
]

/**
 * asyncRoutes 异步路由
 * the routes that need to be dynamically loaded based on user roles   异步路由是根据用户角色动态加载的路由。
 * 每个路由对象包括路径（path）、组件（component）、重定向（redirect）、路由名称（name）、子路由（children）等元素
 */
export const asyncRoutes = [
  {
    path: '/selling',
    component: Layout,
    redirect: '/selling/all',
    name: 'Selling',
    alwaysShow: true,
    meta: {
      title: '销售',
      icon: 'selling'
    },
    children: [{
      path: 'all',
      name: 'SellingAll',
      component: () => import('@/views/selling/all/index'),
      meta: {
        title: '所有销售',
        icon: 'sellingAll'
      }
    },
    {
      path: 'me',
      name: 'SellingMe',
      component: () => import('@/views/selling/me/index'),
      meta: {
        roles: ['editor'],
        title: '我发起的',
        icon: 'sellingMe'
      }
    }, {
      path: 'buy',
      name: 'SellingBuy',
      component: () => import('@/views/selling/buy/index'),
      meta: {
        roles: ['editor'],
        title: '我购买的',
        icon: 'sellingBuy'
      }
    }
    ]
  },
  {
    path: '/donating',
    component: Layout,
    redirect: '/donating/all',
    name: 'Donating',
    alwaysShow: true,
    meta: {
      title: '捐赠',
      icon: 'donating'
    },
    children: [{
      path: 'all',
      name: 'DonatingAll',
      component: () => import('@/views/donating/all/index'),
      meta: {
        title: '所有捐赠',
        icon: 'donatingAll'
      }
    },
    {
      path: 'donor',
      name: 'DonatingDonor',
      component: () => import('@/views/donating/donor/index'),
      meta: {
        roles: ['editor'],
        title: '我发起的捐赠',
        icon: 'donatingDonor'
      }
    }, {
      path: 'grantee',
      name: 'DonatingGrantee',
      component: () => import('@/views/donating/grantee/index'),
      meta: {
        roles: ['editor'],
        title: '我收到的受赠',
        icon: 'donatingGrantee'
      }
    }
    ]
  },
  // {
  //   path: '/addRealestate',
  //   component: Layout,
  //   meta: {
  //     roles: ['admin']
  //   },
  //   children: [{
  //     path: '/addRealestate',
  //     name: 'AddRealestate',
  //     component: () => import('@/views/realestate/add/index'),
  //     meta: {
  //       title: '新增房产',
  //       icon: 'addRealestate'
  //     }
  //   }]
  // },

  {
    path: '/prescription',
    component: Layout,
    redirect: '/prescription/all',
    name: 'Prescription',
    alwaysShow: true,
    meta: {
      title: '病历',
      icon: 'donating'
    },
    children: [
      {
        path: 'all',
        name: 'PrescriptionAll',
        component: () => import('@/views/prescription/list/index'),
        meta: {
          title: '所有病历',
          icon: 'donatingAll'
        }
      },
      {
        path: 'mine',
        name: 'PrescriptionOfMine',
        component: () => import('@/views/prescription/mine/index'),
        meta: {
          // roles: ['patient'],
          title: '我的病历',
          icon: 'donatingDonor'
        }
      },
      {
        path: 'add',
        name: 'Add',
        component: () => import('@/views/prescription/add/index'),
        meta: {
          title: '新增病历',
          icon: 'addRealestate'
        }
      }
    ]
  },

  {
    path: '/insurance',
    component: Layout,
    redirect: '/insurance/all',
    name: 'Insurance',
    alwaysShow: true,
    meta: {
      title: '保险报销',
      icon: 'donating'
    },
    children: [

        {
      path: 'all',
      name: 'InsuranceAll',
      component: () => import('@/views/insuranceCover/list/index'),
      meta: {
        title: '所有的报销记录',
        icon: 'donatingAll'
      }
    },
      {
        path: 'creator',
        name: 'InsuranceCreator',
        component: () => import('@/views/donating/donor/index'),
        meta: {
          roles: ['editor'],
          title: '已发起的报销',
          icon: 'donatingDonor'
        }
      }, {
        path: 'receiver',
        name: 'InsuranceReceiver',
        component: () => import('@/views/donating/grantee/index'),
        meta: {
          roles: ['admin'],
          title: '收到的报销',
          icon: 'donatingGrantee'
        }
      },
      {
        path: 'add',
        name: 'Add',
        component: () => import('@/views/insuranceCover/add/index'),
        meta: {
          title: '新增保险报销',
          icon: 'addRealestate'
        }
      }
    ]
  },

  {
    path: '/drug',
    component: Layout,
    redirect: '/drug/all',
    name: 'Drug',
    alwaysShow: true,
    meta: {
      title: '药品订单',
      icon: 'donating'
    },
    children: [{
      path: 'all',
      name: 'DrugAll',
      component: () => import('@/views/drugOrder/list/index'),
      meta: {
        title: '所有订单',
        icon: 'donatingAll'
      }
    },
      {
        path: '/addDrug',
        name: 'AddDrug',
        component: () => import('@/views/drugOrder/add/index'),
        meta: {
          title: '新增订单',
          icon: 'addRealestate'
        }
      }]
  },

  // 404 page must be placed at the end !!!
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  base: '/',
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
