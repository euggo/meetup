#!/usr/bin/env python3

def content(num):
    # BGN OMIT
    try:
        n = int(num)
    except ValueError:
        return "The value is not an integer."
    except TypeError:
        return "How about no."
    except:
        return "The value is wack, yo."
    # END OMIT

    return "The integer is " + str(n) + "."

def main():
    print(content(".333"))

if __name__ == '__main__':
    main()
