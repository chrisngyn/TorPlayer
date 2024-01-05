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

export function b64toBlob(b64Data: string, contentType?: string, sliceSize?: number): Blob {
  contentType = contentType || "";
  sliceSize = sliceSize || 512;
  const byteCharacters = atob(b64Data);
  const byteArrays = [];

  for (let offset = 0; offset < byteCharacters.length; offset += sliceSize) {
    const slice = byteCharacters.slice(offset, offset + sliceSize);

    const byteNumbers = new Array(slice.length);
    for (let i = 0; i < slice.length; i++) {
      byteNumbers[i] = slice.charCodeAt(i);
    }

    const byteArray = new Uint8Array(byteNumbers);
    byteArrays.push(byteArray);
  }

  return new Blob(byteArrays, { type: contentType });
}

export function arrayBufferToArrayNumber(buffer: ArrayBuffer): Array<number> {
  return Array.from(new Uint8Array(buffer));
}

