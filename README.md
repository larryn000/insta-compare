# Instagram Follow Checker

Find out who doesn't follow you back on Instagram.

## What This Tool Does

This tool compares your Instagram followers and following lists to show you accounts that you follow but who don't follow you back.

## Getting Your Instagram Data

Before using this tool, you need to download your data from Instagram:

1. Open Instagram on your phone or go to [instagram.com](https://instagram.com)
2. Go to **Settings** > **Accounts Center** > **Your information and permissions** > **Download your information**
3. Select **Download or transfer information**
4. Choose your Instagram account
5. Select **Some of your information**
6. Check **Followers and following** only
7. Choose **Download to device**
8. Select **JSON** format (important!)
9. Click **Create files** and wait for Instagram to email you

Once you receive the email (can take a few minutes to a few hours):

1. Download and unzip the file
2. Find these two files:
   - `followers_1.json` (rename to `followers.json`)
   - `following.json`
3. Copy both files to this folder

## Running the Tool

### Option 1: Using the Pre-built Program

If someone gave you a compiled version (`insta-compare` or `insta-compare.exe`):

**On Mac/Linux:**
```
./insta-compare
```

**On Windows:**
```
insta-compare.exe
```

### Option 2: Building From Source

If you have Go installed:

```
go run .
```

## Understanding the Results

The tool will show you a list like this:

```
Users you follow who don't follow you back:
-------------------------------------------
1. someuser123
2. anotheruser
3. thirduser

Total: 3 users
Following: 500 | Followers: 450
```

This means:
- You follow 500 accounts
- 450 accounts follow you
- 3 of the people you follow don't follow you back

## Saving Results to a File

To save the list to a text file:

```
./insta-compare --output results.txt
```

This creates a file called `results.txt` with all the non-followers listed.

## Using Different File Names

If your files have different names or are in a different location:

```
./insta-compare --followers path/to/your/followers.json --following path/to/your/following.json
```

## Frequently Asked Questions

**Q: Is this safe to use?**
A: Yes. This tool runs entirely on your computer. Your data is never uploaded anywhere.

**Q: Will Instagram know I'm using this?**
A: No. This tool only reads the data files you downloaded. It doesn't connect to Instagram.

**Q: Can I unfollow people with this tool?**
A: No. This tool only shows you information. You would need to unfollow people manually through Instagram.

**Q: The numbers don't match my Instagram exactly. Why?**
A: Instagram's data export may be slightly delayed. The numbers should be very close but might not be exact.

**Q: I get an error about "file not found"**
A: Make sure `followers.json` and `following.json` are in the same folder as the program. Check that the files are named exactly right (no extra numbers or spaces).
