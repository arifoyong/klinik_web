import Layout from '../components/Layout'
import {toast} from '../components/Alert/Alert'

export default function Home({sortedPosts}) {

  return (
    <Layout>
      <div className="flex flex-col px-6 py-4 space-y-6">
        Index
      </div>

      <button onClick={() => toast.notify("test message")}
        className="bg-blue-500 text-white px-4 py-2 rounded-lg">
        Test
      </button>

      <button onClick={() => toast.notify("Another error message", "success")}
        className="bg-blue-500 text-white px-4 py-2 rounded-lg">
        Success
      </button>

      <button onClick={() => toast.notify("Another error message", "warning")}
        className="bg-blue-500 text-white px-4 py-2 rounded-lg">
        Warning
      </button>

      <button onClick={() => toast.notify("Another error message", "error")}
        className="bg-blue-500 text-white px-4 py-2 rounded-lg">
        Error
      </button>

    </Layout>
  )
}