import BoxedLayout from 'Container/Boxed';

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
   path: '/boxed',
   component: BoxedLayout,
   redirect: '/boxed/dashboard/ecommerce',
   children: [
      {
         component: Ecommerce,
         path: '/boxed/dashboard/ecommerce',
         meta: {
            requiresAuth: true,
            title: 'message.ecommerce',
             breadcrumb: null
         }
      },
      {
         component: WebAnalytics,
         path: '/boxed/dashboard/web-analytics',
         meta: {
            requiresAuth: true,
            title: 'message.webAnalytics',
             breadcrumb: null
         }
      },
      {
         component: Magazine,
         path: '/boxed/dashboard/magazine',
         meta: {
            requiresAuth: true,
            title: 'message.magazineAndNews',
            breadcrumb: null
         }
      },
      {
         component: News,
         path: '/boxed/dashboard/news',
         meta: {
            requiresAuth: true,
            title: 'message.news',
            breadcrumb: null
         }
      },
      {
         component: Agency,
         path: '/boxed/dashboard/agency',
         meta: {
            requiresAuth: true,
            title: 'message.agency',
             breadcrumb: null
         }
      },
      {
         component: Saas,
         path: '/boxed/dashboard/saas',
         meta: {
            requiresAuth: true,
            title: 'message.saas',
            breadcrumb: null
         }
      },

      
      {
         path: '/boxed/dashboard/crypto',
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
         path: '/boxed/dashboard/crm',
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
