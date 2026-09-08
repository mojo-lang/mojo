// Copyright 2021 Mojo-lang.org
//
// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// [source]()

/// # Time Module 
///

///
type Date {
    year:  Int32 @1
    month: Int32 @2
    day:   Int32 @3
}

///
type DateTime {
    year: Int32 @1
    month: Int32 @2
    day: Int32 @3

    hour: Int32 @4
    minute: Int32 @5
    seconds: Int32 @6
    nanoseconds: Int32 @7

    time_zone: TimeZone @15
}

/// represent timezone information in the locations where different offsets are used in different days of the year
/// or where historical changes have been made to civil time.
type TimeZone {
    /// representing the difference between the local time and UTC.
    /// It must be strictly between -24 * 60 * 60 and 24 * 60 * 60
    offset: Int32 @1 @signed

    /// 
    name:  String @2
}

///
type Timestamp {
    seconds: Int64 @1
    nanoseconds: Int32 @2
}

/// A Duration represents a signed, fixed-length span of time represented as a count of seconds and
/// fractions of seconds at nanosecond resolution. It is independent of any calendar and concepts like "day" or "month".
/// It is related to Timestamp in that the difference between two Timestamp values is a Duration and
/// it can be added or subtracted from a Timestamp. Range is approximately +-10,000 years.
type Duration {
    /// Signed seconds of the span of time. Must be from -315,576,000,000 to +315,576,000,000 inclusive.
    seconds: Int64 @1 @signed

    /// Signed fractions of a second at nanosecond resolution of the span of time. Durations less than one second
    /// are represented with a 0 seconds field and a positive or negative nanoseconds field. For durations of one second or more,
    /// a non-zero value for the nanoseconds field must be of the same sign as the seconds field. Must be
    /// from -999,999,999 to +999,999,999 inclusive.
    nanoseconds: Int32 @2 @signed
}

// Represents a month in the Gregorian calendar.
@case_style("camel")
enum Month {
  // The unspecifed month.
  unspecified @0

  // The month of January.
  january @1

  // The month of February.
  february @2

  // The month of March.
  march @3

  // The month of April.
  april @4

  // The month of May.
  may @5

  // The month of June.
  june @6

  // The month of July.
  july @7

  // The month of August.
  august @8

  // The month of September.
  september @9

  // The month of October.
  october @10

  // The month of November.
  november @11

  // The month of December.
  december @12
}

// Represents a day of week.
@case_style("camel")
enum DayOfWeek {
  // The unspecified day-of-week.
  unspecified @0

  // The day-of-week of Monday.
  monday @1

  // The day-of-week of Tuesday.
  tuesday @2

  // The day-of-week of Wednesday.
  wednesday @3

  // The day-of-week of Thursday.
  thursday @4

  // The day-of-week of Friday.
  friday @5

  // The day-of-week of Saturday.
  saturday @6

  // The day-of-week of Sunday.
  sunday @7
}

// Represents a time of day. The date and time zone are either not significant
// or are specified elsewhere. An API may choose to allow leap seconds. Related
// types are [google.type.Date][google.type.Date] and `google.protobuf.Timestamp`.
type TimeOfDay {
  // Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
  // to allow the value "24:00:00" for scenarios like business closing time.
  hours: Int32 @1

  // Minutes of hour of day. Must be from 0 to 59.
  minutes: Int32 @2

  // Seconds of minutes of the time. Must normally be from 0 to 59. An API may
  // allow the value 60 if it allows leap-seconds.
  seconds: Int32 @3

  // Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
  nanos: Int32 @4
}

///
//func now() -> DateTime

///
//func local_now() -> DateTime

/// ## DateTime
/// 
///
//Timestamp(String)

///
//Timestamp(Date)

///
//func to<>(Timestamp) -> String

///
//func +(Timestamp, Duration) -> DateTime

///
//func -(Timestamp, Duration) -> DateTime

///
//func year(Timestamp) -> Int

///
//func month(Timestamp) -> Int @in([1..12])
