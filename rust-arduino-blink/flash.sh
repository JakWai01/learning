#! /usr/bin/zsh

set -e 

if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    echo "usage: $0 /home/jakobwaibel/Documents/Projects/learning-rust/rust-arduino-blink/target/avr-atmega328p/debug/rust-arduino-blink.elf" >&2
    exit 1
fi

if [ "$#" -lt 1 ]; then
    echo "$0: Expecting a .elf file" >&2
    exit 1
fi

avrdude -q -C/etc/avrdude/avrdude.conf -patmega328p -carduino -P/dev/ttyACM0 -D "-Uflash:w:$1:e"