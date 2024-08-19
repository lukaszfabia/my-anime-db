const MAX_SECONDS = 60;
const MAX_MINUTES = 60 * MAX_SECONDS;
const MAX_HOURS = 24 * MAX_MINUTES;
const MAX_DAYS = 30 * MAX_HOURS;
const MAX_MONTHS = 12 * MAX_DAYS;
const MAX_YEARS = 100 * MAX_MONTHS;
import "date-fns"
import { formatDistanceToNow } from "date-fns";

export default function transformTime(time: Date): string {
    return `${formatDistanceToNow(time)} ago`;
}