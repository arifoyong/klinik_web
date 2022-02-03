import Head from 'next/head'
import Navbar from './NavBar'

const Layout = ({children, ...customMeta}) => {
    const meta = {
        title: "Klinik Auyong",
        description: "Klinik Interface",
        image: "https://eshendetesia.com/images/user-profile.png",
        type: "website",
        ...customMeta
    }

    return (
        <>
            <Head>
                <title>{meta.title}</title>
                <meta name="robots" content="follow, index" />
                <meta name="description" content={meta.description} />

                <meta property="og:type" content={meta.type} />
                <meta property="og:site_name" content="Arif Oyong" />
                <meta property="og:description" content={meta.description} />
                <meta property="og:title" content={meta.title} />
                <meta property="og:image" content={meta.image} />

                {meta.date && (<meta property="article:published_time" content={meta.date} />)}
            </Head>

            
            <div className="flex flex-col max-w-7xl mx-auto">
                <Navbar />
                <div className="">
                    {children}
                </div>
                <div id="toast-container">

                </div>
            </div>
        </>
    )
}

export default Layout