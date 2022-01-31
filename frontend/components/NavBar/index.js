import Link from 'next/link'
import {useTheme} from 'next-themes'
import { useState, useEffect } from 'react'
import { MoonIcon, SunIcon, CubeIcon, MenuIcon } from '@heroicons/react/solid'
import MobileMenu from './MobileMenu'


const NavItem = ({href, text}) => (
    <Link href={href}>
        <a className="hidden sm:inline-block px-3 py-2 
                text-gray-800 dark:text-gray-200
                rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700">
            {text}
        </a>
    </Link>
)

const Navbar = () => {
    const {systemTheme, theme, setTheme} = useTheme()
    const [mounted, setMounted] = useState(false)

    useEffect(() => {
        setMounted(true)
    }, [])


    const renderThemeChanger = () => {
        if (!mounted) return null
        const currentTheme = theme === 'system' ? systemTheme : theme

        if (currentTheme === 'dark') {
            return (
                <SunIcon role="button" onClick={() =>  setTheme('light') }
                            className="w-7 h-7 text-blue-400 darK:text-blue-500 hover:text-blue-800"/>
            )
        } else {
            return (
                <MoonIcon role="button" onClick={() =>  setTheme('dark') }
                            className="w-7 h-7 text-blue-600 hover:text-blue-800 dark:hover:text-blue-500"/>
            )
        }

        
    }

    return (
        <div className="flex items-start justify-between py-4 mb-8">
            <div className="flex items-center space-x-6 w-full">       
                <MobileMenu />   
                <NavItem href="/" text={<CubeIcon className="w-7 h-7 text-blue-600 dark:text-blue-500 hover:text-blue-800"/>}/>
                <NavItem href="/visits" text="Visits"/>
                <NavItem href="/patients" text="Patients"/>
                <NavItem href="/nota" text="Nota"/>
            </div>
            
            {renderThemeChanger()}
            
        </div>
    )
}

export default Navbar