## hugo-figure-maker

This is a script to automatically generate a file with golang tempate
markup, gallery & figure tags populated with its respective keys' values.
To be fed to hugo & photoswipe for a photograpy portfolio site project.


## example output

```golang

{{< gallery >}}

{{< figure link="/img/podejrzane/Kazimierz Dolny(Obrazy podejrzane).jpg"
caption="Kazimierz Dolny." alt="Układanka światła i cienia. Interpretacja dowolna. Elementy „porozrzucane” w kadrze mogą stać się pretekstem do opowiedzenia niejednej historii. Być może z wątkiem kryminalnym?" >}}

what i have is a list of files with jpg files with descriptive names to
be inserted as *caption* key value and matching txt files to be,
optionally, inserted as a *alt* key value.

{{< gallery >}}

```
<figure>

## LICENCE

MIT.
