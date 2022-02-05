import { useState, useEffect } from 'react'
import Layout from '../components/Layout'
import {API} from '../config'



const VisitRegister = () => {
  const [curPatient, setCurPatient] = useState({
    firstname: "",
    lastname: ""
  })
  const [patients, setPatients] = useState(null)

  useEffect(() => {
    const getPatients = async () => {
      const res = await fetch(`${API}/patients`)
      const data = await res.json()
      setPatients(data.data)
    }
    getPatients()
  },[])

  const changeHandler = (e) => {
    setCurPatient({...curPatient, [e.target.name]: e.target.value})
  }


)
  return(
    <Layout>
      <div className="px-2">
        <h2 className="text-3xl py-2">Register Patient Visit</h2>
        <button onClick={() => console.log(patients)}>
          test
        </button>
        <form className="flex-col w-full max-w-sm" >
          <div className="flex items-center py-2">
            <label className="w-1/3" htmlFor="firstname">
              First Name
            </label>
            <input type="text" name="firstname" placeholder="john" value={curPatient.firstname}
                  onChange={changeHandler}
                  className="w-2/3 border-2 border-gray-500 rounded text-gray-800 px-2"
            />
          </div>
          <div className="flex items-center py-2">
            <label className="w-1/3" htmlFor="lastname">
              Last Name
            </label>
            <input type="text" name="lastname" placeholder="doe" value={curPatient.lastname}
                  onChange={changeHandler}
                  className="w-2/3 border-2 border-gray-500 rounded text-gray-800 px-2"
            />
          </div>


          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded-lg">
            Save
          </button>
        </form>
      </div>
    </Layout>
    
  )
}

export default VisitRegister