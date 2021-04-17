import com.google.protobuf.gradle.*

sourceSets {
    main {
        proto {
            srcDir("protos")
            srcDir("grpc")
        }
        java {
            srcDir("build/generated/source/proto/main/grpc")
            srcDir("build/generated/source/proto/main/grpckt")
            srcDir("build/generated/source/proto/main/java")
        }
    }
}

buildscript {
    dependencies {
        classpath("com.google.protobuf:protobuf-gradle-plugin:0.8.15")
    }
}

plugins {
    java
    kotlin("jvm") version "1.4.32"
    id("com.google.protobuf") version "0.8.15"
}

group = "tech.raf924"
version = "1.0-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_11

var grpcVersion = "1.37.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.6.0")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine")
    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("io.grpc:grpc-all:$grpcVersion")
    compileOnly("javax.annotation:javax.annotation-api:1.3.2")
    api("io.grpc:grpc-kotlin-stub:1.0.0")
    implementation("io.grpc:protoc-gen-grpc-kotlin:1.0.0")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.4.3")
    implementation("com.google.protobuf:protobuf-gradle-plugin:0.8.13")
}

tasks.getByName<Test>("test") {
    useJUnitPlatform()
}

tasks.getByName("clean") {
    //delete("src/main/generated")
}

protobuf {
    protoc {
        artifact = "com.google.protobuf:protoc:3.15.8"
    }
    plugins {
        id("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:$grpcVersion"
        }
        id("grpckt") {
            artifact = "io.grpc:protoc-gen-grpc-kotlin:0.1.5"
        }
    }
    generateProtoTasks {
        all().forEach {
            it.plugins {
                id("grpc")
                id("grpckt")
            }
        }
    }
}