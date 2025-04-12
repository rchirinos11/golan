#!/bin/zsh

OUT_DIR="out"
JAR_FILE="App.jar"
MAIN_CLASS="Main"

rm -rf $OUT_DIR

compile() {
  echo "Compiling..."
  mkdir -p $OUT_DIR
  javac -d $OUT_DIR $(find . -name "*.java") || { echo "Compilation failed"; exit 1; }
}

build_jar() {
  echo "Creating Jar..."
  echo "Main-Class: $MAIN_CLASS" > manifest.txt
  jar cfm $JAR_FILE manifest.txt -C $OUT_DIR .
  rm manifest.txt
  echo "Jar file: $JAR_FILE"
}

run_app() {
  echo "Running..."
  java -cp $OUT_DIR $MAIN_CLASS
}

case "$1" in
  jar)
    compile
    build_jar
    ;;
  run)
    compile
    run_app
    ;;
  *)
    echo "Usage: $0 [run | jar]"
    echo "  run     Compile and run application"
    echo "  jar     Build the jar"
    exit 1
    ;;
esac
