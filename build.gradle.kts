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
java.sourceCompatibility = JavaVersion.VERSION_1_8
java.targetCompatibility = JavaVersion.VERSION_1_8


var grpcVersion = "1.37.0"
var protobufVersion = "3.15.8"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
    implementation("javax.annotation:javax.annotation-api:1.3.2")
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.6.0")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine")

    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.4.3")

    api("com.google.protobuf:protobuf-javalite:$protobufVersion")
    api("io.grpc:grpc-kotlin-stub-lite:1.0.0")
}

tasks.getByName<Test>("test") {
    useJUnitPlatform()
}

tasks.configureEach {
    if (name.matches(Regex("^extractInclude([A-Za-z]+)TestProto$"))) {
        enabled = false
    }
}

protobuf {
    protoc {
        artifact = "com.google.protobuf:protoc:$protobufVersion"
    }
    plugins {
        id("java")
    }
    plugins {
        id("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:$grpcVersion"
        }
        id("grpckt") {
            artifact = "io.grpc:protoc-gen-grpc-kotlin:1.0.0:jdk7@jar"
        }
    }
    generateProtoTasks {
        ofSourceSet("main").forEach {
            it.builtins {
                remove("java")
            }
            it.plugins {
                id("java") {
                    option("lite")
                }
                id("grpc"){
                    option("lite")
                }
                id("grpckt"){
                    option("lite")
                }
            }
        }
    }
}