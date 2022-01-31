import  { useState } from 'react'
import Link from 'next/link'
import { MenuIcon, XIcon } from "@heroicons/react/solid"


const MenuList = () => (
    <div className="flex flex-col w-full">
        <Link href="/blogs">
            <a className="px-2 py-2 border-b border-gray-300 dark:border-gray-700
                        hover:bg-gray-200 dark:hover:bg-gray-700">
                Blogs
            </a>
        </Link>
        <Link href="/experiments">
            <a className="px-2 py-2 border-b border-gray-300 dark:border-gray-700
                        hover:bg-gray-200 dark:hover:bg-gray-700">
                Experiments
            </a>
        </Link>
    </div>
)

const MobileMenu = () => {
    const [menuOpen, setMenuOpen] = useState(false)
    
    const ShowMenuIcon = () => (
        <MenuIcon role="button" aria-label="show-menu" className="w-7 h-7"
                onClick={() => setMenuOpen(true)}
        />
    )

    const ShowXIcon = () => (
        <XIcon role="button" aria-label="close-menu" className="w-7 h-7"
                onClick={() => setMenuOpen(false)}
        />
    )

    
    return (
        <div className="inline-block sm:hidden w-full">
            { menuOpen ? ShowXIcon() : ShowMenuIcon() }

            { menuOpen && <MenuList/>}
        </div>
    )
}

export default MobileMenu