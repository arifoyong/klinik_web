

const Patient = ({data}) => (
  <main className="px-4">
    <p className="text-3xl py-4">Patient page</p>
    <table className="table-auto">
      <thead className="bg-blue-700 text-white">
        <tr>
          <td className="px-2 py-2">Firstname</td>
          <td className="px-2 py-2">Lastname</td>
          <td className="px-2 py-2">IC</td>
          <td className="px-2 py-2">DOB</td>
          <td className="px-2 py-2">Email</td>
        </tr>
      </thead>
      <tbody>
        {data.map((patient, idx) => (
          <tr key={idx}>
            <td className="px-2 py-2">{patient.firstname}</td>
            <td className="px-2 py-2">{patient.lastname}</td>
            <td className="px-2 py-2">{patient.ic}</td>
            <td className="px-2 py-2">{patient.dob}</td>
            <td className="px-2 py-2">{patient.email}</td>
          </tr>
          ))}
      </tbody>
    </table>
  </main>
)

export async function getStaticProps(context) {
  const res = await fetch(`http://localhost:8000/patients`)
  const data =await res.json()

  if (!data) {
    return {
      notFound : true
    }
  }

  return {
    props: data
  }
}

export default Patient