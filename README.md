# FigletSampler
Uses figlet as a basis has 350 built in figlet fonts!

Another stupid project to learn Go! This time using file embedding to make a small util that can produce any text you like in ASCII/ANSI using any one of 350 fonts that are built in.

Nope, still no proper testing but at least I'm now aware of it.

After you're initial build, you can simply remove the fonts folder if you like as the build hoovers in all the font files and embeds them.


```
.\FigletSampler.exe -show_random_font -message "Wasting time!"
================================================================================================================
   LOADING : ( 1 ) Small_Slant
=================================================================================================================
 _      __         __  _             __  _           __
| | /| / /__ ____ / /_(_)__  ___ _  / /_(_)_ _  ___ / /
| |/ |/ / _ `(_-</ __/ / _ \/ _ `/ / __/ /  ' \/ -_)_/
|__/|__/\_,_/___/\__/_/_//_/\_, /  \__/_/_/_/_/\__(_)
                           /___/
                           
PS X:\Dev Work\golang\Projects\FigletSampler> .\FigletSampler.exe -show_no_banner -show_random_font -message "Wasting time!"
=================================================================================================================
   LOADING : ( 1 ) Wet_Letter
=================================================================================================================
.-.  .-.  .--.     .---.  _______ ,-..-. .-.  ,--,     _______ ,-.         ,---.   .-.
| |/\| | / /\ \   ( .-._)|__   __||(||  \| |.' .'     |__   __||(||\    /| | .-'   |  )
| /  \ |/ /__\ \ (_) \     )| |   (_)|   | ||  |  __    )| |   (_)|(\  / | | `-.   | /
|  /\  ||  __  | _  \ \   (_) |   | || |\  |\  \ ( _)  (_) |   | |(_)\/  | | .-'   |/
|(/  \ || |  |)|( `-'  )    | |   | || | |)| \  `-) )    | |   | || \  / | |  `--. (
(_)   \||_|  (_) `----'     `-'   `-'/(  (_) )\____/     `-'   `-'| |\/| | /( __.'(_)
                                    (__)    (__)                  '-'  '-'(__)
```
