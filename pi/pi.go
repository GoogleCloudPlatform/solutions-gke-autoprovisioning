// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
)

var (
	calcTime = flag.Duration("calcTime", 10*time.Second, "Calculate pi for this length of time")
	bucket   = flag.String("bucket", "", "Write Pi result to this GCS bucket")
)

func main() {
	flag.Parse()
	val := pi(*calcTime)
	strval := strconv.FormatFloat(val, 'f', -1, 64)
	log.Printf("Calculated Pi for %v: %s\n", *calcTime, strval)
	if *bucket != "" {
		writeToGcs(*bucket, strval)
	}
}

// approximate pi using the Leibniz formula for specified duration
func pi(calcTime time.Duration) float64 {
	f := 0.0
	k := 0.0
	for timeout := time.After(calcTime); ; {
		select {
		case <-timeout:
			return f
		default:
			f += 4 * math.Pow(-1, k) / (2*k + 1)
			k++
		}
	}
}

func writeToGcs(bucketName string, val string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}

	bucket := client.Bucket(bucketName)
	filename := "pi-" + time.Now().UTC().Format(time.RFC3339)
	obj := bucket.Object(filename)

	w := obj.NewWriter(ctx)
	if _, err := fmt.Fprintf(w, "%s", val); err != nil {
		log.Fatalf("Failed to write GCS file: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to write GCS file: %v", err)
	}
	log.Printf("Wrote gs://%s/%s", bucketName, filename)
}
