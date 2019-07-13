# Episode Schedule
Utility lists previous &amp; upcoming TV show episodes.  The shows are currently hardcoded in config.go, and specified via 
their imdb code (found in the show's imdb.com url.)

```console
[user@pc episodes]$ go run *.go 
Episodes:

Loading.. 15.. 14.. 13.. 12.. 11.. 10.. 9.. 8.. 7.. 6.. 5.. 4.. 3.. 2.. 1.. 
Loaded in 967.260021ms

           Elementary  2019.Jul.11   -2d  [s07 e08]  Miss Understood
                       2019.Jul.18    5d  [s07 e09]  On the Scent

      Stranger Things  2019.Jul.04   -9d  [s03 e08]  Chapter Eight: The Battle of Starcourt
                       -- Next episode unknown

                 GLOW  2018.Jun.29 -379d  [s02 e10]  Every Potato Has a Receipt
                       2019.Aug.09   27d  [s03 e01]  Episode 1

  Designated Survivor  2019.Jun.07  -36d  [s03 e10]  #truthorconsequences
                       -- Next episode unknown

          Blue Bloods  2019.May.10  -64d  [s09 e22]  Something Blue
                       2019.Sep.27   76d  [s10 e01]  Episode 1

       The Good Place  2019.Jan.24 -170d  [s03 e13]  Pandemonium
                       2019.Sep.26   75d  [s04 e01]  A Girl from Arizona - Part 1

          The Orville  2019.Apr.25  -79d  [s02 e14]  The Road Not Taken
                       -- Next episode unknown

         Sister Wives  2019.Apr.21  -83d  [s13 e13]  Tell All: Part 2
                       -- Next episode unknown

                Bosch  2019.Apr.19  -85d  [s05 e10]  Creep Signed His Kill
                       -- Next episode unknown

 Star Trek: Discovery  2019.Apr.18  -86d  [s02 e14]  Such Sweet Sorrow, Part 2
                       -- Next episode unknown

            Westworld  2018.Jun.24 -384d  [s02 e10]  The Passenger
                       -- Next episode unknown

              Goliath  2018.Jun.15 -393d  [s02 e08]  Tongue Tied
                       -- Next episode unknown

       Silicon Valley  2018.May.13 -426d  [s05 e08]  Fifty-One Percent
                       -- Next episode unknown

            Mr. Robot  2017.Dec.13 -577d  [s03 e10]  shutdown -r
                       -- Next episode unknown

             Sherlock  2017.Jan.15 -909d  [s04 e03]  The Final Problem
                       -- Next episode unknown

```

## Prerequisites

Uses Go built-ins only.

## About

This is my first golang project.  My goal was simply to get my hands dirty writing some code.
