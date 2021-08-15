package com.advancedJava.teaVM;

import org.teavm.jso.JSBody;

public class main {

    public static void main (String args[]) {

            System.out.println("Hello World!");

            
            log("Jakob");
    }

    @JSBody(params = {"message"}, script = "console.log(message)")
    public static native void log(String message);

    


}
