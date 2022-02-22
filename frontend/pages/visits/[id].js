import {useRouter} from 'next/router'
import Layout from '../../components/Layout'



const Visit = () => {
  const router = useRouter()
  const {id} = router.query
  return (
  <Layout>
      <h2 className="text-3xl py-2">Visits - {id}</h2>
     
    
  </Layout>
  )
}

// export async function getStaticProps(context) {
//   const res = await fetch(`${API}/visits`)
//   const data =await res.json()

//   if (!data) {
//     return {
//       notFound : true
//     }
//   }

//   return {
//     props: data
//   }
// }

export default Visit