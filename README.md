## Quirks:

#### Article folder structure:

Article folder contains folders with a single `article.md` file.
Why not just drop the folders and have each article be an MD file directly in the articles folder?
Because file creation timestamps are not guaranteed based on environment (different OSes and filesystems).

Mod times are guaranteed. The folder's mod time serves as the articles creation timestamp. The mod time of the article file itself, serves as the true mod time for the article.

Mod time of the article folder is used for sorting articles in chronological order.
Mod time of article file is used to reference article edit time.
