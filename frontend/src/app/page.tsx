"use client";

import React, { FC, useEffect, useState } from "react";
import Image from "next/image";
import { motion, AnimatePresence } from "framer-motion";

interface CarouselItemProps {
  id: string;
  src: string;
  alt: string;
  active: boolean;
  children?: React.ReactNode;
}

const CarouselItem: FC<CarouselItemProps> = ({ id, alt, src, active, children }) => {
  return (
    <motion.div
      id={id}
      className={`carousel-item w-full h-full absolute flex items-center justify-center`} // Wyśrodkowanie za pomocą Flexbox
      initial={{ opacity: 0, x: 100 }}
      animate={{ opacity: active ? 1 : 0, x: active ? 0 : -100 }}
      transition={{ duration: 0.5 }}
    >
      <Image
        alt={alt}
        src={src}
        width={1920}
        height={1080}
        priority
        objectFit="cover"
        layout="responsive"
        className="h-full w-full object-cover"
      />
      <div className="absolute text-5xl font-extrabold text-white">
        {children}
      </div>
    </motion.div>
  );
};

export default function Home() {
  const [activeIndex, setActiveIndex] = useState(0);
  const carouselItems = [
    { id: "item1", src: "/images/carousel/jjk.jpg", alt: "carousel item 1" },
    { id: "item2", src: "/images/carousel/kimi_no_na_wa.jpg", alt: "carousel item 2" },
    { id: "item3", src: "/images/carousel/monogatari.jpg", alt: "carousel item 3" },
  ];

  useEffect(() => {
    const interval = setInterval(() => {
      setActiveIndex((prevIndex) => (prevIndex + 1) % carouselItems.length);
    }, 5000);

    return () => clearInterval(interval);
  }, [carouselItems.length]);

  return (
    <main className="py-20">
      <div className="relative overflow-hidden w-full lg:h-[75vh] md:h-[50vh] h-[25vh] mt-1 rounded-t-3xl">
        <AnimatePresence>
          {carouselItems.map((item, index) => (
            <CarouselItem key={item.id} {...item} active={index === activeIndex}>
            </CarouselItem>
          ))}
        </AnimatePresence>
      </div>
      <div className="flex w-full justify-center gap-2 py-2">
        {carouselItems.map((item, index) => (
          <button
            key={item.id}
            onClick={() => setActiveIndex(index)}
            className={`btn btn-xs ${index === activeIndex ? "btn-active" : ""}`}
          >
            {index + 1}
          </button>
        ))}
      </div>
    </main>
  );
}
