// base
import AdminBase from '@/container/AdminBase'

// page
import Dashboard from '@/pages/admin/Dashboard'

export default {
  path: '/admin',
  component: AdminBase,
  redirect: '/admin/main',
  children: [
    {
      path: '/admin',
      name: 'admin-main',
      component: Dashboard
    }
  ]
}
