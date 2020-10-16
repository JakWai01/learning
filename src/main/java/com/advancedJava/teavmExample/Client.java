package com.advancedJava.teavmExample;

import org.teavm.flavour.templates.BindTemplate;
import org.teavm.flavour.widgets.ApplicationTemplate;
import org.apache.commons.codec.binary.Base64;


@BindTemplate("templates/client.html")
//@BindTemplate("templates/test.html")
public class Client extends ApplicationTemplate {
    private String userName = "";
    private String anything = "";
    private static String encodeMe = "{'name':'Jakob', 'age':19}";

    public static void main(String[] args) {
        Client client = new Client();

        client.bind("application-content");
        
        byte[] result1 = client.encodeBase64(encodeMe);
        String result2 = client.decodeBase64(result1);
        System.out.println(result2);
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getAnything() {
        return anything;
    }

    public void setAnything(String anything) {
        this.anything = anything;
    }

    public void setEncodeMe(String encodeMe) {
        Client.encodeMe = encodeMe;
    }

    public String getEncodeMe() {
        return encodeMe;
    }

    public byte[] encodeBase64(String encodeMe) {
        byte[] encodedBytes = Base64.encodeBase64(encodeMe.getBytes());
        return encodedBytes;
    }

    public String decodeBase64(byte[] encodedBytes){
            byte[] decodedBytes = Base64.decodeBase64(encodedBytes);
            return new String(decodedBytes);
            }

    // receive values through html, getters and setters and distribute them via
    // websockets
}
