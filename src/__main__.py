# -*- coding: utf-8 -*-


def get_images_(args):
    from vk import get_images
    get_images.main(args.owner_id)


def vectorize_image_(args):
    import vectorize_image
    vectorize_image.main(args.owner_id)


def download_images_(args):
    from vk import download_images
    download_images.main(args.owner_id)


def main():
    from argparse import ArgumentParser
    parser = ArgumentParser("Images similarity search")
    parser.add_argument("owner_id", type=str, help='Vk owner_id')
    parser.add_argument('-gi', '--get_images', action="store_true", help='scrape images info form the group')
    parser.add_argument('-di', '--download_images', action="store_true", help='download images with info saved in csv')
    parser.add_argument('-vi', '--vectorize_images', action="store_true",
                        help='create a .npz with vector representation of the images')
    args = parser.parse_args()

    if args.get_images:
        get_images_(args)
    if args.download_images:
        download_images_(args)
    if args.vectorize_images:
        vectorize_image_(args)


if __name__ == '__main__':
    main()
