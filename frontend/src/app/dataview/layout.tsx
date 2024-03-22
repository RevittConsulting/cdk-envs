import React from 'react'
import Sidebar from '@/components/datacryp/sidebar'

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className='flex h-full w-full'>
      <Sidebar />
      <div className='w-full overflow-x-hidden'>
        {children}
      </div>
    </div>
  )
}

export default DashboardLayout