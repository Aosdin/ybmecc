import HorizontalLayout from 'Container/HorizontalLayout'

// dashboard components
const Dashboard = () => import('Views/crm/Dashboard');
const Crypto = () => import('Views/crypto/Crypto');
const Ecommerce = () => import('Views/dashboard/Ecommerce');
const WebAnalytics = () => import('Views/dashboard/WebAnalytics');
const Magazine = () => import('Views/dashboard/Magazine');
const News = () => import('Views/dashboard/News');
const Agency = () => import('Views/dashboard/Agency');
const Saas = () => import('Views/dashboard/Saas');

export default {
   path: '/horizontal',
   component: HorizontalLayout,
   redirect: '/horizontal/dashboard/ecommerce',
   children: [
      {
         component: Ecommerce,
         path: '/horizontal/dashboard/ecommerce',
         meta: {
            requiresAuth: true,
            title: 'message.ecommerce',
            breadcrumb: null
         }
      },
      {
         component: WebAnalytics,
         path: '/horizontal/dashboard/web-analytics',
         meta: {
            requiresAuth: true,
            title: 'message.webAnalytics',
            breadcrumb: null
         }
      },
      {
         component: Magazine,
         path: '/horizontal/dashboard/magazine',
         meta: {
            requiresAuth: true,
            title: 'message.magazine',
            breadcrumb: null
         }
      },
      {
         component: News,
         path: '/horizontal/dashboard/news',
         meta: {
            requiresAuth: true,
            title: 'message.news',
            breadcrumb: null
         }
      },
      {
         component: Agency,
         path: '/horizontal/dashboard/agency',
         meta: {
            requiresAuth: true,
            title: 'message.agency',
            breadcrumb: null
         }
      },
      {
         component: Saas,
         path: '/horizontal/dashboard/saas',
         meta: {
            requiresAuth: true,
            title: 'message.saas',
            breadcrumb: null
         }
      },

      
      {
         path: '/horizontal/dashboard/crypto',
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
         path: '/horizontal/dashboard/crm',
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
