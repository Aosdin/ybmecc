import Full from 'Container/Full'

// dashboard components
const Crypto = () => import('Views/crypto/Crypto');
const Ecommerce = () => import('Views/dashboard/Ecommerce');
const WebAnalytics = () => import('Views/dashboard/WebAnalytics');
const Magazine = () => import('Views/dashboard/Magazine');
const News = () => import('Views/dashboard/News');
const Agency = () => import('Views/dashboard/Agency');
const Saas = () => import('Views/dashboard/Saas');
const Dashboard = () => import('Views/crm/Dashboard');


export default {
   path: '/',
   component: Full,
   redirect: '/default/dashboard/ecommerce',
   children: [
      {
         path: '/default/dashboard/ecommerce',
         component: Ecommerce,
         meta: {
            requiresAuth: true,
            title: 'message.ecommerce',
            breadcrumb: null
         }
      },
      {
         path: '/default/dashboard/web-analytics',
         component: WebAnalytics,
         meta: {
            requiresAuth: true,
            title: 'message.webAnalytics',
            breadcrumb: null
         }
      },
      {
         path: '/default/dashboard/magazine',
         component: Magazine,
         meta: {
            requiresAuth: true,
            title: 'message.magazine',
            breadcrumb: null
         }
      },
      {
         path: '/default/dashboard/news',
         component: News,
         meta: {
            requiresAuth: true,
            title: 'message.news',
            breadcrumb: null
         }
      },

      {
         path: '/default/dashboard/agency',
         component: Agency,
         meta: {
            requiresAuth: true,
            title: 'message.agency',
            breadcrumb: null
         }
      },

      {
         component: Saas,
         path: '/default/dashboard/saas',
         meta: {
            requiresAuth: true,
            title: 'message.saas',
            breadcrumb: null
         }
      },

      {
         path: '/default/dashboard/crypto',
         component: Crypto,
         meta: {
            requiresAuth: true,
            title: 'message.crypto',
            breadcrumb: [
              {
                breadcrumbInactive: 'Crypto /'
              },
              {
                breadcrumbActive: 'Crypto'
              }
            ]
         }
      },
      {
         path: '/default/dashboard/crm',
         component: Dashboard,
         meta: {
            requiresAuth: true,
            title: 'message.dashboard',
            breadcrumb: [
              {
                breadcrumbInactive: 'CRM /'
              },
              {
                breadcrumbActive: 'Dashboard'
              }
            ]
         }
      }
   ]
}
