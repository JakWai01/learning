package com.advancedJava.intefaces;

interface Pen {
    String color = "blue";

    void removeCap();
    void write();
    void putCap();
}
public class interfaces implements Pen{
    
    public static void main(String[] args) {
        interfaces pen = new interfaces();
        pen.removeCap();
        pen.write();
        pen.putCap();
        
    }

    public void removeCap() {
        System.out.println("Cap removed!");
    }

    public void write() {
        System.out.println("Write..");
    }

    public void putCap() {
        System.out.println("Cap is back on the "+color+" Pen!");
    }

    
}
