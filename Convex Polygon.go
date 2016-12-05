func isConvex(points [][]int) bool {
    oritations := map[bool]struct{}{}
    sz := len(points)
    for i := 0; i < sz; i += 1 {
        p0, p1, p2 := points[i], points[(i+1)%sz], points[(i+2)%sz]
        if oritation := (p1[0]-p0[0])*(p2[1]-p1[1])-(p1[1]-p0[1])*(p2[0]-p1[0]); oritation != 0 {
            oritations[oritation > 0] = struct{}{}
        }
    }
    return len(oritations) == 1
}
