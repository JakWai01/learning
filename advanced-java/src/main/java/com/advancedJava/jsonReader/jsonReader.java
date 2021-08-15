// package com.advancedJava.jsonReader;

// import java.io.FileReader;
// import java.io.IOException;

// import org.json.simple.JSONObject;
// import org.json.simple.parser.JSONParser;
// import org.json.simple.parser.ParseException;

// public class jsonReader {

//     public static void main(String[] args) throws IOException, ParseException {
//         JSONParser jsonParser = new JSONParser();
    
//     // Parsing the contents of the JSON file
//     JSONObject jsonObject = (JSONObject) jsonParser.parse(new FileReader("src/main/java/com/advancedJava/jsonReader/data.json"));

//     String id = (String)jsonObject.get("ID");
//     String firstName = (String)jsonObject.get("First_Name");
//     String lastName = (String)jsonObject.get("Last_Name");
//     String dateOfBirth = (String)jsonObject.get("Date_Of_Birth");
//     String placeOfBirth = (String)jsonObject.get("Place_Of_Birth");
//     String country = (String)jsonObject.get("Country");

//     System.out.printf("id: %s, firstName: %s, lastName: %s, dateOfBirth: %s, placeOfBirth: %s, country: %s\n", id, firstName, lastName, dateOfBirth, placeOfBirth, country);
//     }
    
// }
