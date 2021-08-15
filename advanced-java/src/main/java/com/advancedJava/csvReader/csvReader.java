package com.advancedJava.csvReader;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class csvReader {

    public static void main(String[] args) throws IOException {
        
        
       List<List<String>> records = new ArrayList<>();
       try (BufferedReader br = new BufferedReader(new FileReader("src/main/java/com/advancedJava/csvReader/data.csv"))) {
           String line;
           while ((line = br.readLine()) !=  null) {
               String[] values = line.split(",");
               records.add(Arrays.asList(values));
           }
       }

       System.out.println(records.toString());
    }
}
