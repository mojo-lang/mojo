import csv

type Location {
    test: String!
    str:  Int32
}

Csv<Location>('test.csv') | {
                              test: "od" + $.test
                              str: $.str + 33
                            }
                          | Csv<Location>('new.csv')