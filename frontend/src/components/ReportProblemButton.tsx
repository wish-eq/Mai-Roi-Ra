import React, { useState } from "react";
import ReportPopup from "@/components/ReportPopup";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";

const ReportProblemButton = () => {
  const [showModal, setShowModal] = useState(false);

  return (
    <div className="lg:pt-8 pt-6 px-20 text-black">
      <div className="flex flex-row justify-center w-full mt-[30px] mb-[50px]">
        <button
          className="border-slate-400 border flex justify-center flex-row items-center rounded-full 
                lg:h-[40px] md:h-[35px] h-[35px] 
                lg:w-[140px] md:w-[110px] w-[110px] hover:scale-105 duration-300 py-[10px] px-[10px]
                lg:text-[17px] md:text-[11px] text-[11px]"
          onClick={(e) => {
            e.stopPropagation();
            e.preventDefault();
            setShowModal(true);
          }}
        >
          <span className="mr-[5px]">
            <AddCircleOutlineIcon />
          </span>{" "}
          Report
        </button>
      </div>

      <ReportPopup isVisible={showModal} onClose={() => setShowModal(false)} />
    </div>
  );
};

export default ReportProblemButton;
