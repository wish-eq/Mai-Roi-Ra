'use client'
import styles from "@/styles/FontPage.module.css"
import { useRouter } from "next/navigation";
import { useState } from "react";
import SuccessModal from "./SuccessModal";
import PhotoSizeSelectActualIcon from '@mui/icons-material/PhotoSizeSelectActual';

const CreateEventForm = () => {
    const router = useRouter();
    const [showModal,setShowModal] = useState(true);
    const [eventName,setEventName] = useState("");
    const [eventDescription,setEventDescription] = useState("");
    const [uploading, setUploading] = useState(false);
    const [selectedImage, setSelectedImage] = useState("");
    const [selectedFile, setSelectedFile] = useState<File>();

    const handleSubmit = () => {
        setShowModal(true);
    }

    const uploadImage = () => {

    }

    return (
        <div className={`${styles.Roboto} w-full h-full pt-[5%] pl-[5%] pr-[10%]`}>
            <form onSubmit={handleSubmit} action="" className="w-full h-full">
                
                {/* Form */}
                <div className="flex flex-row flex-wrap justify-between h-[80%] w-full">

                    {/* Left Form */}
                    <div className="w-[60%] h-[90%] space-y-[10%]">
                        <div className="w-full h-[13%] relative">
                            <input className="border-[1px] border-gray-300 w-[60%] indent-4 h-full text-[18px]
                            rounded-md"
                            type="text" placeholder="Event Name" 
                            value={eventName} onChange={(e) => setEventName(e.target.value)}/>

                            {eventName.length != 0 && (
                                <div className="absolute top-[-8px] px-2 left-2 bg-white transition-all text-xs text-gray-400">
                                    Event Name
                                </div>
                            )}
                        </div>

                        <div className="w-full h-[13%] relative">
                            <input className="border-[1px] border-gray-300 w-[60%] indent-4 h-full text-[18px]
                            rounded-md"
                            type="text" placeholder="Date Start/End"/>

                            {eventName.length != 0 && (
                                <div className="absolute top-[-8px] px-2 left-2 bg-white transition-all text-xs text-gray-400">
                                    Date Start/End
                                </div>
                            )}
                        </div>

                        <div className="w-full h-[45%] relative">
                            <textarea className="border-[1px] border-gray-300 w-full indent-4 h-full text-[18px]
                            rounded-md py-[10px]"
                            placeholder="Event Description"
                            value={eventDescription} onChange={(e)=>setEventDescription(e.target.value)}/>

                            {eventDescription.length != 0 && (
                                <div className="absolute top-[-8px] px-2 left-2 bg-white transition-all text-xs text-gray-400">
                                    Event Name
                                </div>
                            )}
                        </div> 
                    </div>

                    {/* Right Form */}
                    <div className="h-[90%] w-[30%] border-[1px] border-gray-300 rounded-md flex justify-center items-center
                    cursor-pointer">
                        <div className="text-center text-gray-400 hover:text-black"
                            onClick={uploadImage}>
                            <PhotoSizeSelectActualIcon className="w-[15%] h-[15%]"/>
                            <div className="">
                                Add Picture
                            </div>
                        </div>
                    </div>

                </div>

                {/* Button */}
                <div className="w-full flex flex-row flex-wrap justify-around">
                    <div className="w-[30%]">
                        <button className="bg-red-500 py-[17px] px-[100px] rounded-full text-white"
                        onClick={() => router.push("/homepage")}>
                            Cancel
                        </button>
                    </div>
                    
                    <div className="w-[30%]">
                        <button type="submit" className="bg-blue-400 py-[17px] px-[110px] rounded-full text-white">
                            Done
                        </button>
                    </div>
                </div>
                <SuccessModal topic="Evented Created" isVisible={showModal} onClose={()=>setShowModal(false)}/>
            </form>
        </div>
    )
}

export default CreateEventForm;