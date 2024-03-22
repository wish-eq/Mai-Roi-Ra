'use client'
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

export default function UserSupportAndService() {
  const pathname = usePathname();

  console.log(pathname);

  return (
    <div className="bg-white text-black h-full">
      <div className="lg:mr-24 border-r border-b bg-white">
        <div className="w-full text-2xl pt-20">
          {["FAQ", "Problem"].map((tabName) => (
            <Link href={`/supportandservice/${tabName.toLowerCase()}`}
              key={tabName}
              className={`relative px-8 py-2 overflow-hidden`}
            >
              <span
                className={`transition duration-500 ease-in-out ${
                  pathname === `supportandservice/${tabName.toLowerCase()}`
                    ? "text-yellow-500"
                    : "text-gray-800 hover:text-gray-500"
                }
                `}
              >
                {tabName}
              </span>
              <span
                className={`absolute inset-x-0 bottom-0 h-0.5 bg-yellow-500 transition-all duration-500 ease-out ${
                  pathname === `supportandservice/${tabName.toLowerCase()}` ? "scale-x-100" : "scale-x-0"
                }`}
                style={{
                  transformOrigin: "center",
                }}
              ></span>
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
}
