// base
import AdminBase from '@/container/AdminBase'

// page
import Dashboard from '@/pages/admin/Dashboard'
import Signin from '@/pages/admin/Signin'
import Teacher from '@/pages/admin/Teacher'
import Student from '@/pages/admin/Student'

export default {
  path: '/admin',
  component: AdminBase,
  children: [
    {
      path: '/admin',
      name: 'admin-main',
      component: Dashboard
    },
    {
      path: 'signin',
      name: 'admin-signin',
      component: Signin
    },
    {
      path: 'teacher',
      name: 'admin-teacher',
      component: Teacher
    },
    {
      path: 'student',
      name: 'admin-student',
      component: Student
    }
  ]
}
