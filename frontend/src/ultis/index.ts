
export function bytesLengthToSize(length: number): string {
  const sizes = ["B", "KB", "MB", "GB", "TB"];
  if (length == 0) return "0 B";
  const i = Math.floor(Math.log(length) / Math.log(1024));
  return (length / Math.pow(1024, i)).toFixed(2) + " " + sizes[i];
}

export function isVideoFile(fileName: string): boolean {
  const extension = getFileExtension(fileName);
  const videoExtensions = ["mp4", "mov", "avi", "mkv", "mpg", "wmv"];
  return videoExtensions.includes(extension);
}

export function getFileExtension(fileName: string): string {
  const parts = fileName.split(".");
  return parts[parts.length - 1].toLowerCase();
}
