# Episode Schedule
Utility lists previous &amp; upcoming TV show episodes.  The shows are currently hardcoded in config.go, and specified via 
their imdb code (found in the show's imdb.com url.)

```console
[user@pc episodes]$ go run *.go 
Episodes:  [loaded in 1.062331s]

           Elementary  2019.Jul.11   -3d  [s07 e08]  Miss Understood
                       2019.Jul.18    4d  [s07 e09]  On the Scent

      Stranger Things  2019.Jul.04  -10d  [s03 e08]  Chapter Eight: The Battle of Starcourt
                       -- Next episode unknown

                 GLOW  2018.Jun.29 -380d  [s02 e10]  Every Potato Has a Receipt
                       2019.Aug.09   26d  [s03 e01]  Episode 1

  Designated Survivor  2019.Jun.07  -37d  [s03 e10]  #truthorconsequences
                       -- Next episode unknown

          Blue Bloods  2019.May.10  -65d  [s09 e22]  Something Blue
                       2019.Sep.27   75d  [s10 e01]  Episode 1

       The Good Place  2019.Jan.24 -171d  [s03 e13]  Pandemonium
                       2019.Sep.26   74d  [s04 e01]  A Girl from Arizona - Part 1

          The Orville  2019.Apr.25  -80d  [s02 e14]  The Road Not Taken
                       -- Next episode unknown

         Sister Wives  2019.Apr.21  -84d  [s13 e13]  Tell All: Part 2
                       -- Next episode unknown

                Bosch  2019.Apr.19  -86d  [s05 e10]  Creep Signed His Kill
                       -- Next episode unknown

 Star Trek: Discovery  2019.Apr.18  -87d  [s02 e14]  Such Sweet Sorrow, Part 2
                       -- Next episode unknown

            Westworld  2018.Jun.24 -385d  [s02 e10]  The Passenger
                       -- Next episode unknown

              Goliath  2018.Jun.15 -394d  [s02 e08]  Tongue Tied
                       -- Next episode unknown

       Silicon Valley  2018.May.13 -427d  [s05 e08]  Fifty-One Percent
                       -- Next episode unknown

            Mr. Robot  2017.Dec.13 -578d  [s03 e10]  shutdown -r
                       -- Next episode unknown

             Sherlock  2017.Jan.15 -910d  [s04 e03]  The Final Problem
                       -- Next episode unknown
n

```

## Prerequisites

Uses Go built-ins only.

## About

This is my first golang project.  My goal was simply to get my hands dirty writing some code.
