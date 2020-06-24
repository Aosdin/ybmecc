import MiniSidebarLayout from 'Container/MiniSidebarLayout'

// dashboard components
const Ecommerce = () => import('Views/dashboard/Ecommerce');
const WebAnalytics = () => import('Views/dashboard/WebAnalytics');
const Magazine = () => import('Views/dashboard/Magazine');
const News = () => import('Views/dashboard/News');
const Agency = () => import('Views/dashboard/Agency');
const Saas = () => import('Views/dashboard/Saas');
const Crypto = () => import('Views/crypto/Crypto');
const Dashboard = () => import('Views/crm/Dashboard');

export default {
   path: '/mini',
   component: MiniSidebarLayout,
   redirect: '/mini/dashboard/ecommerce',
   children: [
      {
         component: Ecommerce,
         path: '/mini/dashboard/ecommerce',
         meta: {
            requiresAuth: true,
            title: 'message.ecommerce',
            breadcrumb: null
         }
      },
      {
         component: WebAnalytics,
         path: '/mini/dashboard/web-analytics',
         meta: {
            requiresAuth: true,
            title: 'message.webAnalytics',
            breadcrumb: null
         }
      },
      {
         component: Magazine,
         path: '/mini/dashboard/magazine',
         meta: {
            requiresAuth: true,
            title: 'message.magazine',
            breadcrumb: null
         }
      },
      {
         component: News,
         path: '/mini/dashboard/news',
         meta: {
            requiresAuth: true,
            title: 'message.news',
            breadcrumb: null
         }
      },
      {
         component: Agency,
         path: '/mini/dashboard/agency',
         meta: {
            requiresAuth: true,
            title: 'message.agency',
            breadcrumb: null
         }
      },
      {
         component: Saas,
         path: '/mini/dashboard/saas',
         meta: {
            requiresAuth: true,
            title: 'message.saas',
            breadcrumb: null
         }
      },
      
      {
         path: '/mini/dashboard/crypto',
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
         path: '/mini/dashboard/crm',
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
