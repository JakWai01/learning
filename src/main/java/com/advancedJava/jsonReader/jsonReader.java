package com.advancedJava.jsonReader;

import java.io.FileReader;
import java.io.IOException;

import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

public class jsonReader {

    public static void main(String[] args) throws IOException, ParseException {
        JSONParser jsonParser = new JSONParser();
    
    // Parsing the contents of the JSON file
    JSONObject jsonObject = (JSONObject) jsonParser.parse(new FileReader("src/main/java/com/advancedJava/jsonReader/data.json"));

    String value = (String)jsonObject.get("ID");

    System.out.println(value);
    }
    
}
