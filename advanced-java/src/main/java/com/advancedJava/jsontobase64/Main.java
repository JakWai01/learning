package com.advancedJava.jsontobase64;

import org.apache.commons.codec.binary.Base64;

public class Main {
    
    public static void main(String[] args) {

        Main asd = new Main();

        byte[] result1 = asd.encodeBase64("{'name':'Jakob', 'age':19}");
        String result2 = asd.decodeBase64(result1);
        System.out.println(result2);
    }

    public byte[] encodeBase64(String encodeMe) {
        byte[] encodedBytes = Base64.encodeBase64(encodeMe.getBytes());
        return encodedBytes;
    }

    public String decodeBase64(byte[] encodedBytes){
            byte[] decodedBytes = Base64.decodeBase64(encodedBytes);
            return new String(decodedBytes);
            }
}
