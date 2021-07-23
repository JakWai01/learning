pub fn structs() {
    #[derive(Debug)]
    struct Student {
        name: String,
        mail: String,
        matriculation_number: u8,
    }

    let jakob = Student {
        name: String::from("Jakob Waibel"),
        mail: String::from("jakob@university.com"),
        matriculation_number: 1,
    };

    impl std::fmt::Display for Student {
        fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
            write!(
                f,
                "Student (name: {}, mail: {}, matriculation_number: {})",
                self.name, self.mail, self.matriculation_number
            )
        }
    }

    println!("{}", jakob);
}
