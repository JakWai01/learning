package com.advancedJava.teaVMwasm;

import org.teavm.interop.Export;

public class Main {
    
    @Export(name = "thePurposeOfLife")
    public static int getPurposeOfLife() {
        return 42;
    }
}
