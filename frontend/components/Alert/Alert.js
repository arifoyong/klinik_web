import  { render, unmountComponentAtNode }  from "react-dom"
import { XIcon } from "@heroicons/react/solid"

const Toast = (props) => {
  let toastClass = ""
  switch (props.type) {
    case "success":
      toastClass = "flex justify-between bg-green-500 text-white px-4 py-2 rounded-xl"
      break
    case "warning":
      toastClass = "flex justify-between bg-yellow-500 text-white px-4 py-2 rounded-xl"
      break
    case "error":
      toastClass = "flex justify-between bg-red-500 text-white px-4 py-2 rounded-xl"
      break
    default:
      toastClass = "flex justify-between bg-gray-500 text-white px-4 py-2 rounded-xl"
  }

  return (
    <div className="fixed bottom-4 right-4 ">
        <div className={toastClass}>
          <p> {props.message}</p>
          <XIcon onClick={props.close} 
                className="w-4 h-4 ml-4 hover:cursor-pointer"/>
        </div>
 
    </div>
  )
}


export const toast = {
  remove: () => {
    unmountComponentAtNode(document.getElementById('toast-container'))
    toast.currentToast = false
    if (toast.timeout) {
      clearTimeout(toast.timeout)
      toast.timeout = null
    }
   },
   currentToast: false,
  timeout: null,
  notify: (message, type="info") => {
    let duration = 5

    if(toast.currentToast) { 
      toast.remove()
    }

    render(<Toast 
      message={message} 
      slideIn={true} 
      type={ type}
      close = {() => toast.remove()}
      duration={duration} />, document.getElementById('toast-container'));
   

    toast.currentToast = true
    toast.timeout = setTimeout( toast.remove, duration*1000)
  }
}