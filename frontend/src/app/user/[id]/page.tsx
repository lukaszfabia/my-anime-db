"use client";

import { useParams } from "next/navigation"

export default function User() {
    const { id } = useParams();
    return (
        <div>
            User <span>{id}</span>
        </div>
    )
}