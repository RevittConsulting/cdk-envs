import React from 'react'
import Sidebar from '@/components/sidebar'

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <div className='md:h-full md:z-[80] md:fixed md:w-60 md:flex md:flex-col hidden'>
        <Sidebar />
      </div>
      <main className='md:pl-60'>
      {children}
      </main>
    </>
  )
}

export default DashboardLayout