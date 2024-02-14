'use server'
import { revalidatePath, revalidateTag } from "next/cache"
import { redirect } from "next/navigation"
import UpdateEvent from "@/libs/updateEvent"


export async function HandleUpdateEvent(id:string,name:string,activity:string, dateRange:string, price:number, location:string, district:string, 
    province:string, description:string, imageSrc:string){

    try {
        const res = await UpdateEvent(id,name,activity,location,district,province,price,description,
            imageSrc,dateRange.substring(0,10),dateRange.substring(13,23))
        console.log(res)
        console.log("Update Event successful")
    } catch (err) {
        console.log("Error during update event: ", err)
    }
    revalidateTag(`events`);
    revalidateTag('event');
    revalidatePath(`/homepage`);
    revalidatePath('/profile');
    revalidatePath(`/event/${id}`);
    redirect(`/profile`);
}


