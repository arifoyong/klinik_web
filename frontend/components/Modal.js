import React, {useEffect, useRef, useState} from 'react'
import ReactDOM from 'react-dom'
import { XCircleIcon } from '@heroicons/react/solid'


const Modal = ({show, onClose, children, title}) => {
  const [isBrowser, setIsBrowser] = useState(false)

  useEffect(() => {
    setIsBrowser(true)
  }, [])

  const handleCloseClick = (e) => {
    e.preventDefault()
    onClose()
  }
  
  const ModalContent = show ? (
    <div className="fixed z-10 inset-0 overflow-y-auto">
      <div className="flex items-center justify-center min-h-screen 
                      pt-4 px-4 pb-20 text-center sm:block sm:p-0 bg-gray-500 bg-opacity-75">
        <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="false">&#8203;</span>
        <div className="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div className="flex justify-between">
              <p className="text-xl font-bold mb-4">{title}</p>
              <XCircleIcon onClick={handleCloseClick}
                  className="w-6 h-6 hover:cursor-pointer"/>
            </div>
            
            {children}
          </div>
        </div>
      </div>

      {/* <p className="text-blue-800">Modal</p>
        <button onClick={handleCloseClick}>close</button> */}
    </div>
  ) : null;

  if (isBrowser) {
    return ReactDOM.createPortal(
      ModalContent,
      document.getElementById("modal-root")
    )
  } else {
    return null
  }


}

export default Modal