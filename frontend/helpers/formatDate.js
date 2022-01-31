export function dd_mmm_yyyy(strDate) {
  const monthNames =["Jan","Feb","Mar","Apr",
                      "May","Jun","Jul","Aug",
                      "Sep", "Oct","Nov","Dec"];
  const dt = new Date(strDate)

  const day = dt.getDate()
  const mth = monthNames[dt.getMonth()]
  const yr = dt.getFullYear()

  return `${day}-${mth}-${yr}`
}

export function yyyy_mm_dd(strDate) {
  const dt = new Date(strDate)
  return dt.toISOString().split('T')[0]
}


export function getToday() {
  const today = new Date();
  return today.toISOString().split('T')[0]
}
