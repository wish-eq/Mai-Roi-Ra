"use client";
import React, { useState } from "react";
import AdminProblemList from "./AdminProblemList";
import ProblemList from "./ProblemList";

interface Props {
  datas: any[];
}

export default function AdminSupportAndService({ datas }: Props) {
  const [activeTab, setActiveTab] = useState("Problem");

  const handleTabClick = (tab: string) => {
    setActiveTab(tab);
  };

  return (
    <div className="bg-white text-black h-full">
      <div className="lg:mr-24 border-r border-b bg-white">
        <div className="w-full text-2xl pt-20">
          {["Problem", "Replied"].map((tabName) => (
            <button
              key={tabName}
              className={`relative px-8 py-2 overflow-hidden`}
              onClick={() => handleTabClick(tabName)}
            >
              <span
                className={`transition duration-500 ease-in-out ${
                  activeTab === tabName
                    ? "text-yellow-500"
                    : "text-gray-800 hover:text-gray-500"
                }`}
              >
                {tabName}
              </span>
              <span
                className={`absolute inset-x-0 bottom-0 h-0.5 bg-yellow-500 transition-all duration-500 ease-out ${
                  activeTab === tabName ? "scale-x-100" : "scale-x-0"
                }`}
                style={{
                  transformOrigin: "center",
                }}
              ></span>
            </button>
          ))}
        </div>
      </div>
      {activeTab == "Problem" && (
        <div className="my-8 px-4 lg:px-10">
          <AdminProblemList datas={datas}></AdminProblemList>
        </div>
      )}
    </div>
  );
}
