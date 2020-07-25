// base
import AdminBase from '@/container/AdminBase'

// page
import Dashboard from '@/pages/admin/Dashboard'
import Teacher from '@/pages/admin/Teacher'

export default {
  path: '/admin',
  component: AdminBase,
  redirect: '/admin/main',
  children: [
    {
      path: '/admin',
      name: 'admin-main',
      component: Dashboard
    },
    {
      path: '/teacher',
      name: 'admin-teacher',
      component: Teacher
    }
  ]
}
