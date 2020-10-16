package com.advancedJava.app;

import java.util.Scanner;

/**
 * Main App
 */
public class Application {

	public static void main(String[] args) {
		System.out.println("Hello World!");

		int a = 2;
		int result = 0;

		Pen pen = new Pen("blue");
		System.out.println(pen.getColor());

		Scanner scan = new Scanner(System.in);
		int iterations = scan.nextInt();

		for (int i = 0; i < iterations; i++) {
			result += a;
		}

		System.out.println(result); 
		scan.close();

		Felttip felt = new Felttip("red");
		System.out.println(felt.getColor());
}
}
