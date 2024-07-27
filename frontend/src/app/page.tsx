'use client';

import { useEffect } from "react";
import api from "../lib/api";

export default function Home() {

  useEffect(() => {
    api.post("/login", { username: "test", password: "test" }).then((res) => {
      console.log(res.data);
    }).catch((err) => {
      console.log(err);
    });



    api.get("/auth/account/me").then((res) => {
      console.log(res.data);
    }).catch((err) => {
      console.log(err);
    });
  }, []);


  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">

    </main>
  );
}
